/**
 * Coordinates pending sources and button ownership for one management list.
 * Record-specific copying and presentation remain in the calling controller.
 * @module components/management-draft-presentation/coordinator
 * @author GPT-6.1 Sol
 */
import type { CloneOperation } from "./operation";

/**
 * Keeps replacement operations isolated from stale presentation completions.
 * Create one instance per list and register only after synchronous copying succeeds.
 * @example
 * const coordinator = new CloneCoordinator();
 * coordinator.cancel();
 * // Copy the draft synchronously, then create its CloneOperation.
 * coordinator.register(source, button, operation);
 * // In the presentation's finally block, after disposing the operation:
 * coordinator.finish(button, operation);
 * @author GPT-6.1 Sol
 */
export class CloneCoordinator {
    private active: CloneOperation | undefined;
    private pendingSource: HTMLElement | undefined;
    private readonly buttonOwners = new WeakMap<HTMLButtonElement, CloneOperation>();

    /**
     * Reports whether a source already has a non-aborted presentation in progress.
     * @example `if(coordinator.isPending(source)) { return; }`
     * @author GPT-6.1 Sol
     */
    isPending(source: HTMLElement): boolean {
        return this.active !== undefined && !this.active.controller.signal.aborted && this.pendingSource === source;
    }

    /**
     * Aborts the active presentation before replacement without releasing newer ownership.
     * @example `coordinator.cancel(); // Before synchronously copying another draft.`
     * @author GPT-6.1 Sol
     */
    cancel(): void {
        this.active?.controller.abort();
    }

    /**
     * Associates a successfully copied draft's presentation with its source and button.
     * @example `coordinator.register(source, button, operation);`
     * @author GPT-6.1 Sol
     */
    register(source: HTMLElement, button: HTMLButtonElement, operation: CloneOperation): void {
        this.active = operation;
        this.pendingSource = source;
        this.buttonOwners.set(button, operation);
    }

    /**
     * Releases busy and pending state only while the completing operation still owns it.
     * Call from presentation cleanup, including after cancellation or presentation failure.
     * @example `coordinator.finish(button, operation);`
     * @author GPT-6.1 Sol
     */
    finish(button: HTMLButtonElement, operation: CloneOperation): void {
        if(this.buttonOwners.get(button) === operation) {
            button.removeAttribute("aria-busy");
            this.buttonOwners.delete(button);
        }

        if(this.active === operation) {
            this.active = undefined;
            this.pendingSource = undefined;
        }
    }
}
