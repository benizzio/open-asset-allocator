#!/usr/bin/env sh
# Runs containerized Playwright E2E modes with isolated Docker Compose resources.
# Usage: ./e2e.sh {local|ci|ui|headed|report|logs|clean} [Playwright arguments...]
# Set E2E_NO_CACHE=1 for a clean build or E2E_PULL=1 to refresh base images.
# Authored by: OpenCode

set -u

CDPATH=''
script_dir=$(cd -P "$(dirname "$0")" && pwd)
project_root=$script_dir
docker_dir="$project_root/src/main/docker"
artifact_dir="$project_root/target/e2e-results"
base_compose_file="$docker_dir/docker-compose-e2e.yml"
local_compose_file="$docker_dir/docker-compose-e2e-local.yml"
ci_compose_file="$docker_dir/docker-compose-e2e-ci.yml"
project_name=${COMPOSE_PROJECT_NAME:-open-asset-allocator-e2e}
cleanup_required=0
cleanup_complete=0

fail() {
  # Prints an actionable error before returning a nonzero status.
  printf '%s\n' "Error: $*" >&2
  return 1
}

require_dependencies() {
  # Verifies the only host tools required to run the E2E environment.
  command -v docker >/dev/null 2>&1 || {
    fail "Docker is required."
    return 1
  }
  docker compose version >/dev/null 2>&1 || {
    fail "Docker Compose is required."
    return 1
  }
  command -v make >/dev/null 2>&1 || {
    fail "Make is required."
    return 1
  }
}

validate_build_options() {
  # Validates opt-in Docker build controls before selecting images.
  case "${E2E_NO_CACHE:-0}" in
    0|1) ;;
    *) fail "E2E_NO_CACHE must be 0 or 1."; return 1 ;;
  esac
  case "${E2E_PULL:-0}" in
    0|1) ;;
    *) fail "E2E_PULL must be 0 or 1."; return 1 ;;
  esac
}

validate_project_name() {
  # Prevents E2E lifecycle commands from addressing production or development stacks.
  case "$project_name" in
    open-asset-allocator-e2e|open-asset-allocator-e2e-*) ;;
    *) fail "COMPOSE_PROJECT_NAME must be open-asset-allocator-e2e or start with open-asset-allocator-e2e-." ;;
  esac
}

prepare_environment() {
  # Exports host identity for artifact ownership and creates the artifact destination.
  E2E_UID=${E2E_UID:-$(id -u)}
  E2E_GID=${E2E_GID:-$(id -g)}
  export COMPOSE_PROJECT_NAME="$project_name"
  export E2E_UID E2E_GID

  mkdir -p "$artifact_dir" || return 1
  [ -w "$artifact_dir" ] || fail "E2E artifact directory is not writable: $artifact_dir"
}

compose() {
  # Runs Compose against the shared E2E configuration and a selected topology overlay.
  compose_file=$1
  shift
  docker compose --project-name "$project_name" --file "$base_compose_file" --file "$compose_file" "$@"
}

compose_debug() {
  # Runs the local E2E configuration with the headed-debug profile enabled.
  docker compose --project-name "$project_name" --file "$base_compose_file" --file "$local_compose_file" --profile debug "$@"
}

compose_all() {
  # Addresses all E2E service definitions for safe logs and explicit cleanup.
  docker compose --project-name "$project_name" --file "$base_compose_file" --file "$local_compose_file" --file "$ci_compose_file" --profile debug "$@"
}

build_images() {
  # Builds selected services while allowing Docker to resolve cache invalidation.
  build_compose_file=$1
  build_profile=$2
  shift 2

  if [ "${E2E_NO_CACHE:-0}" -eq 1 ]; then
    set -- --no-cache "$@"
  fi
  if [ "${E2E_PULL:-0}" -eq 1 ]; then
    set -- --pull "$@"
  fi

  if [ "$build_profile" = debug ]; then
    compose_debug build "$@"
  else
    compose "$build_compose_file" build "$@"
  fi
}

capture_logs() {
  # Saves Compose service logs while containers still exist for failure diagnosis.
  compose "$selected_compose_file" logs --no-color --timestamps >"$artifact_dir/compose.log" 2>&1
}

cleanup_stack() {
  # Captures diagnostics and removes only the selected E2E project resources once.
  [ "$cleanup_required" -eq 1 ] || return 0
  [ "$cleanup_complete" -eq 0 ] || return 0

  cleanup_complete=1
  cleanup_status=0
  capture_logs || cleanup_status=1
  compose "$selected_compose_file" down --volumes --remove-orphans || cleanup_status=1
  return "$cleanup_status"
}

on_exit() {
  # Preserves the command status while reporting log-capture or cleanup failures.
  exit_status=$?

  if ! cleanup_stack && [ "$exit_status" -eq 0 ]; then
    exit_status=1
  fi

  trap - 0
  exit "$exit_status"
}

on_signal() {
  # Converts handled termination signals to their conventional exit statuses.
  exit "$1"
}

start_stack() {
  # Clears stale resources, builds selected images, and starts the requested topology.
  compose "$selected_compose_file" down --volumes --remove-orphans || return 1
  cleanup_required=1

  case "$mode" in
    ci)
      build_images "$ci_compose_file" standard monolith e2e-runner || return 1
      compose "$selected_compose_file" up --detach --remove-orphans db migration-engine monolith || return 1
      ;;
    headed)
      build_images "$local_compose_file" debug backend frontend e2e-debug || return 1
      compose_debug up --detach --remove-orphans db migration-engine backend frontend || return 1
      ;;
    *)
      build_images "$local_compose_file" standard backend frontend e2e-runner || return 1
      compose "$selected_compose_file" up --detach --remove-orphans db migration-engine backend frontend || return 1
      ;;
  esac
}

print_local_url() {
  # Prints a loopback URL using the configured local port.
  url_port=$1
  url_host=${E2E_HOST_IP:-127.0.0.1}

  if [ "$url_host" = "127.0.0.1" ]; then
    url_host=localhost
  fi

  printf 'http://%s:%s\n' "$url_host" "$url_port"
}

run_tests() {
  # Runs Playwright in the selected topology and forwards unmodified CLI filters.
  case "$mode" in
    ui)
      printf 'Application: %s\n' "$(print_local_url "${E2E_FRONTEND_PORT:-8082}")"
      printf 'Playwright UI: %s\n' "$(print_local_url "${E2E_UI_PORT:-9324}")"
      compose "$selected_compose_file" run --rm --no-deps --publish "${E2E_HOST_IP:-127.0.0.1}:${E2E_UI_PORT:-9324}:9324" e2e-runner npm test -- --ui --ui-host=0.0.0.0 --ui-port=9324 "$@"
      ;;
    headed)
      printf 'Application: %s\n' "$(print_local_url "${E2E_FRONTEND_PORT:-8082}")"
      printf 'noVNC: %s\n' "$(print_local_url "${E2E_NOVNC_PORT:-7900}")"
      compose_debug run --rm --no-deps --service-ports e2e-debug npm test -- --headed "$@"
      ;;
    *)
      compose "$selected_compose_file" run --rm --no-deps e2e-runner npm test -- "$@"
      ;;
  esac
}

serve_report() {
  # Serves the most recent HTML report without starting application services.
  report_dir="$artifact_dir/html-report"
  [ -f "$report_dir/index.html" ] || fail "No HTML report found at $report_dir. Run an E2E suite first."

  selected_compose_file=$local_compose_file
  build_images "$local_compose_file" standard e2e-runner || return 1
  cleanup_required=1
  printf 'HTML report: %s\n' "$(print_local_url "${E2E_REPORT_PORT:-9323}")"
  compose "$selected_compose_file" run --rm --no-deps --publish "${E2E_HOST_IP:-127.0.0.1}:${E2E_REPORT_PORT:-9323}:9323" e2e-runner npm exec -- playwright show-report /workspace/target/e2e-results/html-report --host=0.0.0.0 --port=9323
}

main() {
  # Selects the requested lifecycle mode and returns its command status.
  mode=${1:-}
  shift || true

  case "$mode" in
    local|ci|ui|headed|report|logs|clean) ;;
    *)
      printf '%s\n' "Usage: $0 {local|ci|ui|headed|report|logs|clean} [Playwright arguments...]" >&2
      return 2
      ;;
  esac

  require_dependencies || return 1
  validate_build_options || return 1
  validate_project_name || return 1
  prepare_environment || return 1

  case "$mode" in
    clean)
      compose_all down --volumes --remove-orphans
      ;;
    logs)
      compose_all logs --follow --no-color --timestamps
      ;;
    report)
      serve_report
      ;;
    *)
      case "$mode" in
        ci) selected_compose_file=$ci_compose_file ;;
        *) selected_compose_file=$local_compose_file ;;
      esac
      start_stack || return 1
      run_tests "$@"
      ;;
  esac
}

trap on_exit 0
trap 'on_signal 129' HUP
trap 'on_signal 130' INT
trap 'on_signal 143' TERM

main "$@"
