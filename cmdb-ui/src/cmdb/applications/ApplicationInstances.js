import React from "react";
import { connect } from 'react-redux';

// from core components
import EnvironmentSelects from "../environments/EnvironmentSelects";
import AppButton from "./AppButton";

// from Redux
import { listApplicationInstances } from "../../actions";

// Table format is something like:
// <Table
//   tableHeaderColor="primary"
//   tableHead={["Name", "Country", "City", "Salary"]}
//   tableData={[
//     ["Dakota Rice", "Niger", "Oud-Turnhout", "$36,738"],
//     ["Minerva Hooper", "Curaçao", "Sinaai-Waas", "$23,789"],
//   ]}
// />

class ApplicationInstances extends React.Component {
  constructor(props) {
    super(props);
    this.state = { envName: ""};
  }

  selectEnvironment = (envId, envName) => {
    this.props.listApplicationInstances(envId);
    this.setState({envName} );
  }

  // applicationInstanceTableData = () => {
  //   const appInstanceTableData = this.props.applicationInstances.filter( (applicationInstance) => {
  //     if (applicationInstance.name && applicationInstance.name.length ) {
  //       return true;
  //     } else {
  //       return false;
  //     }
  //   }).map( (applicationInstance) => {
  //     return [applicationInstance.name];
  //   });
  //
  //   return appInstanceTableData;
  // };
  //
  // render() {
  //   return(
  //     <>
  //       < EnvironmentSelects
  //         envName={this.state.envName}
  //         selectEnvironment={this.selectEnvironment}
  //       />
  //       <Table
  //         tableHeaderColor="primary"
  //         tableHead={["Name"]}
  //         tableData={this.applicationInstanceTableData()}
  //       />
  //     </>
  //
  //   );
  // }

  renderAppInstances = () => {
    return(
      this.props.applicationInstances.filter( (applicationInstance) => {
        if (applicationInstance.name && applicationInstance.name.length ) {
          return true;
        } else {
          return false;
        }
      }).map( (applicationInstance) => {
        return <AppButton name={applicationInstance.name}/>;
      }));
  }
  render() {
    return(
      <>
        < EnvironmentSelects
          envName={this.state.envName}
          selectEnvironment={this.selectEnvironment}
        />
        {this.renderAppInstances()}
      </>
    );
  }
}

const mapStateToProps = state => {
  return {
    applicationInstances: Object.values(state.applicationInstances),
  };
};

export default connect(mapStateToProps, {listApplicationInstances})(ApplicationInstances);
