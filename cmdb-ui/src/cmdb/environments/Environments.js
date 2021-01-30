import React from "react";
import { connect } from 'react-redux';

// from core components
import Table from "../../components/Table/Table";

// from Redux
import { listEnvironments } from "../../actions";

// Table format is something like:
// <Table
//   tableHeaderColor="primary"
//   tableHead={["Name", "Country", "City", "Salary"]}
//   tableData={[
//     ["Dakota Rice", "Niger", "Oud-Turnhout", "$36,738"],
//     ["Minerva Hooper", "Curaçao", "Sinaai-Waas", "$23,789"],
//   ]}
// />

class Environments extends React.Component {
  componentDidMount() {
    this.props.listEnvironments();
  }

   environmentTableData = () => {
    const envTableData = this.props.environments.filter( (environment) => {
      if (environment.name && environment.name.length ) {
        return true;
      } else {
        return false;
      }
    }).map( (environment) => {
        return [environment.name];
    });

    return envTableData;
  };

  render() {
    return(
      <>
        <Table
          tableHeaderColor="primary"
          tableHead={["Name"]}
          tableData={this.environmentTableData()}
        />
      </>

    );
  }
}

const mapStateToProps = state => {
  return {
    environments: Object.values(state.environments)
  };
};

export default connect(mapStateToProps, { listEnvironments })(Environments);
