import {Application} from "./application.model";

export class ChartVersion {
    constructor(public id: string,
                public repo: string,
                public version: string,
                public application: Application[]){}
}
