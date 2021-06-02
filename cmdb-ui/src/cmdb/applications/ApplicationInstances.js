import React from "react";
import { connect } from 'react-redux';

// from core components
import EnvironmentSelects from "../environments/EnvironmentSelects";
import AppInstanceButton from "./AppInstanceButton";

// from Redux
import IconButton from "@material-ui/core/IconButton";
import Cancel from "@material-ui/icons/Cancel";
import { makeStyles } from "@material-ui/core/styles";
import styles from "../../assets/jss/material-dashboard-react/views/dashboardStyle";
import {withStyles} from "@material-ui/core";

import { listApplicationInstances,
  selectEnvironment,
  createManifest,
  clearManifest,
  createValues,
  clearValues } from "../../actions";

const useStyles = makeStyles(styles);

class ApplicationInstances extends React.Component {

  constructor(props) {
    super(props);
    this.state = {showApp: false, showValue: false};
  }

  selectEnvironment = (envId) => {
    this.props.selectEnvironment(envId);
    this.props.listApplicationInstances({envId});
  }

  renderAppInstances = () => {
    const enabledAppInstances = this.props.applicationInstances.filter((applicationInstance) => {
      return (applicationInstance.name && applicationInstance.name.length && applicationInstance.enable)
    });
    return(
      <>
        <h3>{enabledAppInstances.length} Application Instances</h3>
        { enabledAppInstances.map( (applicationInstance) => {
          return <AppInstanceButton key={applicationInstance.id} app={applicationInstance} onClick={(showValues) => {this.showAppView(showValues, applicationInstance)}}/>;
        })}
      </>
    );
  }

  closeAppView = () => {
    this.setState({showApp: false});
    this.props.clearManifest("");
    this.props.clearValues("");
  }

  showAppView = (showValues, applicationInstance) => {
    this.setState({showApp: true, showValues: showValues, applicationInstance: applicationInstance});
    if (applicationInstance.enable) {
      if (showValues) {
        this.props.createValues(applicationInstance.id);
      } else {
        this.props.createManifest(applicationInstance.id);
      }
    } else {
      this.props.clearManifest("App Instance is disabled!");
      this.props.clearValues("App Instance is disabled!");
    }
  }

  renderAppInstance = () => {
    return(
      <div>
        <IconButton onClick={this.closeAppView}>
          <Cancel />
        </IconButton>
        <h3>{this.state.applicationInstance.name}</h3>
        <div>{this.props.manifest.status}</div>
        <div>{(this.state.showValues) ? this.props.manifest.values : this.props.manifest.artifact}</div>
      </div>
    );
  }

  render() {
    return(
      <>
        < EnvironmentSelects
          envId={this.props.envId}
          selectEnvironment={this.selectEnvironment}
        />
        {(this.state.showApp) ? this.renderAppInstance() : this.renderAppInstances()}
      </>
    );
  }
}

const mapStateToProps = state => {
  return {
    envId: state.selectedEnvId,
    applicationInstances: Object.values(state.applicationInstances),
    manifest: state.manifest,
    values: state.values,
  };
};

export default connect(mapStateToProps,
  {listApplicationInstances,
    createManifest,
    clearManifest,
    createValues,
    clearValues,
    selectEnvironment,})
(withStyles(useStyles)(ApplicationInstances));
