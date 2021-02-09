import cmdb, { headers } from "../api/cmdb";
import {apiListApplicationInstances} from "../api/applicationInstances";
import { LIST_APPLICATIONS,
  LIST_APPLICATION_INSTANCES,
  LIST_ENVIRONMENTS, 
  LIST_LIFECYCLES,
  LIST_CHART_VERSIONS,
  SELECT_ENVIRONMENT, } from "./types";

export const listApplications = () => async dispatch => {
  const response = await cmdb.get('/v1/applications?_order_by=name&_fields=id,name', {headers});
  dispatch({type: LIST_APPLICATIONS, payload: response.data.results});
}

export const listApplicationInstances = ({ envId = "", appId = "" }) => async dispatch => {
  const response = await apiListApplicationInstances(envId, appId);
  dispatch({type: LIST_APPLICATION_INSTANCES, payload: response.data.results});
}

export const listEnvironments = () => async dispatch => {
  const response = await cmdb.get('/v1/environments?_order_by=name&_fields=id,name,lifecycle_id', {headers});
  dispatch({type: LIST_ENVIRONMENTS, payload: response.data.results});
}

export const listLifecycles = () => async dispatch => {
  const response = await cmdb.get('/v1/lifecycles?_order_by=name&_fields=id,name', {headers});
  dispatch({type: LIST_LIFECYCLES, payload: response.data.results});
}

export const listChartVersions = () => async dispatch => {
  const response = await cmdb.get('/v1/chart_versions?_order_by=name&_fields=id,name,repo,version', {headers});
  dispatch({type: LIST_CHART_VERSIONS, payload: response.data.results});
}

export const selectEnvironment = ( envId = "" ) =>  {
  return({type: SELECT_ENVIRONMENT, payload: envId } );
}
