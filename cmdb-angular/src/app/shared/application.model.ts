import {ApplicationInstance} from "./application_instance.model";

export class Application {
    constructor(public id: string,
                public name: string,
                public description: string,
                public repo: string,
                public ApplicationInstances: ApplicationInstance[]){}
}
