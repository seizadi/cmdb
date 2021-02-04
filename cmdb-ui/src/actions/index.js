import cmdb from "../api/cmdb";
import { LIST_APPLICATIONS,
  LIST_APPLICATION_INSTANCES,
  LIST_ENVIRONMENTS, 
  LIST_LIFECYCLES,
  LIST_CHART_VERSIONS,
  SELECT_ENVIRONMENT, } from "./types";

const headers = {
  'Content-Type': 'application/json',
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SUQiOjF9.GsXyFDDARjXe1t9DPo2LIBKHEal3O7t3vLI3edA7dGU',
};

export const listApplications = () => async dispatch => {
  const response = await cmdb.get('/v1/applications?_order_by=name&_fields=id,name', {headers});
  dispatch({type: LIST_APPLICATIONS, payload: response.data.results});
}

export const listApplicationInstances = ({ envId = "", appId = "" }) => async dispatch => {
  let url = '/v1/application_instances';
  let search = "";
  if (envId.length > 0 && appId.length > 0) {
    search = '&_filter=environment_id=="' + encodeURIComponent(envId) + " AND " +
      'application_id=="' + encodeURIComponent(appId) +'"';
  } else if (envId.length > 0) {
    search = '&_filter=environment_id=="' + encodeURIComponent(envId) + '"';
  } else if (appId.length > 0 ) {
    search = '&_filter=application_id=="' + encodeURIComponent(appId) + '"';
  }
  url = url +
    '?_order_by=name&_fields=id,name,application_id,environment_id,chart_version_id' +
    search;
  const response = await cmdb.get( url, {headers});
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
