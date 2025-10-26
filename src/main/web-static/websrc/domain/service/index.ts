import { AllocationDomainService } from "./allocation-service";
import { mapToPortfolio } from "./mapping/portfolio-mapping";
import { mapToSerializableCompleteAllocationPlans } from "./mapping/allocation-plan/serializable-complete-allocation-plan-mapping";
import { mapToCompleteAllocationPlan } from "./mapping/allocation-plan/complete-allocation-plan-mapping";

export const DomainService = {

    allocation: AllocationDomainService,

    mapping: { mapToPortfolio, mapToCompleteAllocationPlan, mapToSerializableCompleteAllocationPlans },
};