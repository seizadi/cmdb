import cmdb from "../api/cmdb";
import { LIST_APPLICATIONS, LIST_ENVIRONMENTS, LIST_LIFECYCLES } from "./types";

const headers = {
  'Content-Type': 'application/json',
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SUQiOjF9.GsXyFDDARjXe1t9DPo2LIBKHEal3O7t3vLI3edA7dGU',
};

export const listApplications = () => async dispatch => {
  const response = await cmdb.get('/v1/applications?_order_by=name&_fields=id,name', {headers});
  dispatch({type: LIST_APPLICATIONS, payload: response.data.results});
}

export const listEnvironments = () => async dispatch => {
  const response = await cmdb.get('/v1/environments?_order_by=name&_fields=id,name,lifecycle_id', {headers});
  dispatch({type: LIST_ENVIRONMENTS, payload: response.data.results});
}

export const listLifecycles = () => async dispatch => {
  const response = await cmdb.get('/v1/lifecycles?_order_by=name&_fields=id,name', {headers});
  dispatch({type: LIST_LIFECYCLES, payload: response.data.results});
}
