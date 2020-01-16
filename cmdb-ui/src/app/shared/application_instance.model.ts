import {ChartVersion} from "./chart_version.model";

export class ApplicationInstance {
    constructor(public id: string,
                public config_yaml: string,
                public chart: ChartVersion){}
}
