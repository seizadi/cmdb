import React from "react";
import { connect } from 'react-redux';

import { listApplications } from "../../actions";
import Table from "../../components/Table/Table";

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
  componentDidMount() {
    this.props.listApplications();
  }

   applicationTableData = () => {
    const appTableData = this.props.applications.filter( (application) => {
      if (application.name && application.name.length ) {
        return true;
      } else {
        return false;
      }
    }).map( (application) => {
        return [application.name];
    });

    return appTableData;
  };

  render() {
    return(
      <>
        <Table
          tableHeaderColor="primary"
          tableHead={["Name"]}
          tableData={this.applicationTableData()}
        />
      </>

    );
  }
}

const mapStateToProps = state => {
  return {
    applications: Object.values(state.applications)
  };
};

export default connect(mapStateToProps, { listApplications })(Applications);
