export interface Sheet {
    id: number;
    title: string;
    description: string;
    owner: string;
    protection: number;
    searchable: boolean;
    exe_times: number;
    comments: number;
    evaluations: number;
    star: number;
    created_date: Date;
    mod_date: Date;
}
