import { AllocationStructure, AllocationStructureDTO } from "./allocation";

export function mapAllocationStructure(allocationStructureDTO: AllocationStructureDTO): AllocationStructure {
    return {
        hierarchy: allocationStructureDTO.hierarchy.map((hierarchyLevelDTO, index) => ({
            ...hierarchyLevelDTO,
            index: index,
        })),
    };
}