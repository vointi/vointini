/* Do not change, this code is generated from Golang structs */


export interface DTOOK {
    msg: string;
}
export interface DTONewId {
    id: number;
}
export interface DTOEntry {
    id: number;
    activity: string;
    description: string;
    achievement: number;
    levels: {[key: string]: number};
    tags: string[];
}
export interface DTOEntryAdd {
    activity: string;
    description: string;
    achievement: number;
    levels: {[key: string]: number};
    tags: string[];
}
export interface DTOEntryLevel {
    id: number;
    name: string;
    key: string;
    show: boolean;
    worst: string;
    added_at: string;
}
export interface DTOEntryLevelUpdate {
    name: string;
    key: string;
    show: boolean;
    worst: string;
}
export interface DTOTag {
    id: number;
    added_at: string;
    name: string;
    shortname: string;
}
export interface DTOTask {
    id: number;
    added_at: string;
    completed_at?: string;
    title: string;
    description: string;
    refid?: number;
}
export interface DTOTimer {
    id: number;
    s: number;
    f: string;
}
export interface DTOTimerAdd {
    title: string;
    seconds: number;
}
export interface DTOReoccurringTaskAdd {
    title: string;
    s: number;
}
export interface DTOReoccurringTask {
    id: number;
    added_at: string;
    title: string;
}
export interface DTOWeight {
    weight: number;
    added: string;
    id: number;
}
export interface DTOWeightAdd {
    weight: number;
}
export interface DTOHeight {
    height: number;
    added: string;
    id: number;
}
export interface DTOHeightAdd {
    height: number;
}
export interface DTOResolutionFile {
    id: number;
    added_at: string;
    name: string;
}
export interface DTOResolutions {
    id: number;
    entityid: number;
    added_at: string;
    name: string;
    decisiondate: string;
    sentdate: string;
    startdate: string;
    enddate: string;
    files: DTOResolutionFile[];
}
export interface DTOResolutionsUpdate {
    entityid: number;
    name: string;
    decisiondate: string;
    sentdate: string;
    startdate: string;
    enddate: string;
}
export interface DTOResolutionEntity {
    id: number;
    name: string;
}

export interface DTOTestMADRSAnswers {
    a1: number;
    a2: number;
    a3: number;
    a4: number;
    a5: number;
    a6: number;
    a7: number;
    a8: number;
    a9: number;
    a10: number;
}