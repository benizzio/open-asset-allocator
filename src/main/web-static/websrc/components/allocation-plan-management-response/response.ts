/**
 * Maps paginated allocation-plan API records into the management template context.
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */

import { AllocationPlanDTO, SerializableCompleteAllocationPlan } from "../../domain/allocation-plan";
import { AllocationPlanType } from "../../domain/allocation";
import { DomainService } from "../../domain/service";
import { PortfolioPage } from "../../pages";
import { getBrowserLocalCalendarDate } from "../../utils/calendar-date";
import { type BeforeSwapEventDetail } from "../../infra/htmx";
import paginatedList, { PaginationMode } from "../paginated-list";

/** Maps the current plan page and adds the synthetic new-plan form only to an initial/reset page.
 *
 * @param responseBody - Paginated JSON returned by the allocation-plan endpoint.
 * @param eventDetail - HTMX response event used to validate and stage the page cursor.
 * @returns Serialized management template context for the accepted plan records.
 * @example Register this function as the GET response transform for the allocation-plan management list.
 * @author GPT-6 Luna
 */
export function mapForAllocationPlanManagement(
    responseBody: string,
    eventDetail: BeforeSwapEventDetail,
): string {
    return paginatedList.transformPageResponse<AllocationPlanDTO, ReturnType<typeof mapPlanPage>>(
        responseBody,
        eventDetail,
        plan => plan.id,
        (plans, mode) => mapPlanPage(plans, mode),
    );
}

/** Maps one accepted plan page to the existing allocation-plan management template context.
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
function mapPlanPage(plans: AllocationPlanDTO[], mode: PaginationMode) {
    const portfolioDTO = PortfolioPage.getContextPortfolio();

    const completeAllocationPlanSet = DomainService.mapping.mapToSerializablePortfolioCompleteAllocationPlanSet(
        portfolioDTO,
        plans,
    );

    if(mode === PaginationMode.Initial) {
        const hierarchy = DomainService.mapping.mapToPortfolio(portfolioDTO).allocationStructure.hierarchy;
        const topHierarchyLevel = hierarchy.at(-1);

        const planZero: SerializableCompleteAllocationPlan = {
            allocationPlan: {
                id: 0,
                name: "",
                type: AllocationPlanType.ASSET_ALLOCATION_PLAN,
                plannedDate: getBrowserLocalCalendarDate(),
                details: [],
            },
            fractalHierarchicalPlan: { subLevel: topHierarchyLevel, topAllocations: [] },
            topLevelKey: topHierarchyLevel.name,
        };

        completeAllocationPlanSet.completeAllocationPlans.unshift(planZero);
    }

    return completeAllocationPlanSet;
}
