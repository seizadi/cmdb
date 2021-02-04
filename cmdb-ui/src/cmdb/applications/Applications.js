import React from "react";
import { connect } from 'react-redux';

// from core components
import AppButton from "./AppButton";

// from Redux
import { listApplications,
  listApplicationInstances,
  listEnvironments,
  listChartVersions,
  listLifecycles } from "../../actions";
import AppGraph from "./AppGraph";
import IconButton from "@material-ui/core/Button";
import CancelIcon from "@material-ui/icons/Cancel";
import { createStyles, withStyles } from '@material-ui/core';

const styles = (theme) => createStyles({
  root: {
    '& > *': {
      margin: theme.spacing(1),
    },
  }
});

// Table format is something like:
// <Table
//   tableHeaderColor="primary"
//   tableHead={["Name", "Country", "City", "Salary"]}
//   tableData={[
//     ["Dakota Rice", "Niger", "Oud-Turnhout", "$36,738"],
//     ["Minerva Hooper", "Curaçao", "Sinaai-Waas", "$23,789"],
//   ]}
// />

class Applications extends React.Component {
  constructor() {
    super();
    this.state = { showGraph: false };
  }

  componentDidMount() {
    this.props.listApplications();
    this.props.listEnvironments();
    this.props.listLifecycles();
    this.props.listChartVersions();
  }

  //  applicationTableData = () => {
  //   const appTableData = this.props.applications.filter( (application) => {
  //     if (application.name && application.name.length ) {
  //       return true;
  //     } else {
  //       return false;
  //     }
  //   }).map( (application) => {
  //       return [application.name];
  //   });
  //
  //   return appTableData;
  // };
  //
  // render() {
  //   return(
  //     <>
  //       <Table
  //         tableHeaderColor="primary"
  //         tableHead={["Name"]}
  //         tableData={this.applicationTableData()}
  //       />
  //     </>
  //
  //   );
  // }

  handleAppClick = (application) => {
    this.props.listApplicationInstances({appId: application.id});
    this.setState( { showGraph: true });
  }

  handleAppGraphCancel = () => {
    this.setState( { showGraph: false });
  }

  renderApps = () => {
    return (
      this.props.applications.filter( (application) => {
        return (application.name && application.name.length);
      }).map( (application) => {
        return < AppButton key={application.id}
                           app={application}
                           onClick={() => {this.handleAppClick(application)}} />
      }));
  };

  render() {
    if (this.state.showGraph) {
      return (
        <>
          <div className={this.props.classes.root}>
            <IconButton aria-label="cancel" onClick={this.handleAppGraphCancel}>
              <CancelIcon />
            </IconButton>
          </div>
          <AppGraph applicationInstances={this.props.applicationInstances}
                    environments={this.props.environments}
                    lifecycles={this.props.lifecycles}
                    chartVersions={this.props.chartVersions}
          />
        </>
      );
    }
    return( this.renderApps() );
  }
}

const mapStateToProps = state => {
  return {
    applications: Object.values(state.applications),
    applicationInstances: Object.values(state.applicationInstances),
    environments: state.environments,
    chartVersions: state.chartVersions,
    lifecycles: state.lifecycles,
  };
};

export default connect( mapStateToProps,
  {listApplications,
    listApplicationInstances,
    listEnvironments,
    listChartVersions,
    listLifecycles })
(withStyles(styles)(Applications));
