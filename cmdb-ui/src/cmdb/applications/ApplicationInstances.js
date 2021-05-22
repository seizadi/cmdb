import React from "react";
import { connect } from 'react-redux';

// from core components
import EnvironmentSelects from "../environments/EnvironmentSelects";
import AppButton from "./AppButton";

// from Redux
import IconButton from "@material-ui/core/IconButton";
import Cancel from "@material-ui/icons/Cancel";
import { makeStyles } from "@material-ui/core/styles";
import styles from "../../assets/jss/material-dashboard-react/views/dashboardStyle";
import {withStyles} from "@material-ui/core";

import { listApplicationInstances, selectEnvironment, createManifest, clearManifest } from "../../actions";

const useStyles = makeStyles(styles);

class ApplicationInstances extends React.Component {

  constructor(props) {
    super(props);
    this.state = {showApp: false};
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
          return <AppButton key={applicationInstance.id} app={applicationInstance} onClick={() => {this.showAppView(applicationInstance)}}/>;
        })}
      </>
    );
  }

  closeAppView = () => {
    this.setState({showApp: false});
    this.props.clearManifest("");
  }

  showAppView = (applicationInstance) => {
    this.setState({showApp: true, applicationInstance: applicationInstance});
    if (applicationInstance.enable) {
      this.props.createManifest(applicationInstance.id);
    } else {
      this.props.clearManifest("App Instance is disabled!");
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
        <div>{this.props.manifest.artifact}</div>
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
  };
};

export default connect(mapStateToProps,
  {listApplicationInstances,
    createManifest,
    clearManifest,
    selectEnvironment,})
(withStyles(useStyles)(ApplicationInstances));
