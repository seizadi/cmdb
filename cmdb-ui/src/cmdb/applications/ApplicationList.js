import React from "react";
import { connect } from 'react-redux';

import { listApplications } from "../../actions";

class ApplicationList extends React.Component {
  componentDidMount() {
    console.log('call list application...')
    this.props.listApplications();
  }

  renderList() {
    return this.props.applications.map( (application) => {
      return (
        <div className="item" key={ application.id } >
          <div className="content">
              { application.name }
          </div>
        </div>
      );
    });
  }

  render() {
    return(
      <div>
        <h2>Applications</h2>
        <div className="ui celled list">
          { this.renderList() }
        </div>
      </div>

    );
  }
}

const mapStateToProps = state => {
  return {
    applications: Object.values(state.applications)
  };
};

export default connect(mapStateToProps, { listApplications })(ApplicationList);
