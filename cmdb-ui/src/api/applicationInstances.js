import cmdb, {headers} from './cmdb'

export const apiListApplicationInstances = (envId = "", appId = "" ) => {
  let url = '/cmdb/v1/application_instances';
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
    '?_order_by=name&_fields=id,name,enable, application_id,environment_id,chart_version_id,config_yaml' +
    search;

  return cmdb.get( url, {headers});
}

export const apiGetManifest = (appInstanceId = "" ) => {
  let url = '/v1/manifest';
  return cmdb.post( url, { app_instance_id: appInstanceId }, {headers} );
}
