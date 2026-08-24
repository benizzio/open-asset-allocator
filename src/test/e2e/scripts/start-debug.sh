#!/usr/bin/env sh
# Starts the display services required for headed Playwright execution and noVNC.
# Usage: start-debug.sh npm test -- --headed
# Authored by: OpenCode
set -u

display_process_id=""
vnc_process_id=""
novnc_process_id=""

cleanup() {
  for process_id in "$novnc_process_id" "$vnc_process_id" "$display_process_id"; do
    if [ -n "$process_id" ]; then
      kill "$process_id" 2>/dev/null || true
    fi
  done

  for process_id in "$novnc_process_id" "$vnc_process_id" "$display_process_id"; do
    if [ -n "$process_id" ]; then
      wait "$process_id" 2>/dev/null || true
    fi
  done
}

trap cleanup EXIT
trap 'exit 143' INT TERM

export DISPLAY="${DISPLAY:-:99}"
display_number="${DISPLAY#:}"
display_number="${display_number%%.*}"
display_socket="/tmp/.X11-unix/X${display_number}"

Xvfb "$DISPLAY" -screen 0 1440x900x24 -nolisten tcp &
display_process_id=$!

attempt=0
while [ ! -S "$display_socket" ] && [ "$attempt" -lt 30 ]; do
  sleep 1
  attempt=$((attempt + 1))
done

if [ ! -S "$display_socket" ]; then
  exit 1
fi

x11vnc -display "$DISPLAY" -forever -shared -nopw -rfbport 5900 &
vnc_process_id=$!

/usr/share/novnc/utils/novnc_proxy --vnc localhost:5900 --listen 7900 &
novnc_process_id=$!

"$@"
exit_code=$?
exit "$exit_code"
