import React from "react";
import { Graph } from 'react-d3-graph';
import { connect } from 'react-redux';

class AppGraph extends React.Component {

  constructor(props) {
    super(props);


    const mynodes = {
      nodes: [
        {id: 'dev', color: 'red', x: 100, y: 200, rank: 0},
        {id: 'integration', color: 'red', x: 250, y: 200, rank: 1},
        {id: 'stage', color: 'red', x: 400, y: 200, rank: 2},
        {id: 'prod', color: 'red', x: 550, y: 200, rank: 3},
        {id: 'gov', color: 'red', x: 700, y: 200, rank: 4},
      ]
    };

    const mylinks = {
      links: [
        {source: 'dev', target: 'integration'},
        {source: 'integration', target: 'stage'},
        {source: 'stage', target: 'prod'},
        {source: 'prod', target: 'gov'},
      ]
    };

    // The graph configuration, you only need to pass down properties
    //that you want to override, otherwise default ones will be used
    const myConfig = {
      // height: 400,
      // width: 800,
      // automaticRearrangeAfterDropNode: false,
      directed: true,
      nodeHighlightBehavior: true,
      disableLinkForceConfig: false, // --> new API config
      node: {
        color: 'lightgreen',
        size: 120,
        highlightStrokeColor: 'blue'
      },
      link: {
        highlightColor: 'lightblue'
      }
    };

    this.state = {data: {nodes: mynodes.nodes, links: mylinks.links}, config: myConfig};
  }

  lookupLifeCycleName(name) {
    switch (name) {
      case 'gov-dev':
      case 'gov-stage':
      case 'gov-prod':
        return 'gov';
      default:
        return name;
    }
  }

  componentDidMount() {
    let dataSnapShot = this.state.data;
    // Add the charts to Graph
    let chartMap = {}
    this.props.applicationInstances.forEach((appInstance) => {
      if (this.props.graph.showDisabled ||  appInstance.enable) {
        chartMap[appInstance.chart_version_id] = appInstance.chart_version_id;
      }
    });
    const chartIds = Object.values(chartMap)
    chartIds.forEach(id => {
      const chartName = this.props.chartVersions[id].repo + ':' + this.props.chartVersions[id].version;
      if (chartName && chartName.length > 0) {
        dataSnapShot.nodes.push({id: chartName, color: 'gray'});
      };
    });

    this.props.applicationInstances.forEach((appInstance) => {
      if ( this.props.graph.showDisabled || appInstance.enable) {
        const env = this.props.environments[appInstance.environment_id];
        const lifecycle = this.props.lifecycles[env.lifecycle_id];
        const lifecycleName = this.lookupLifeCycleName(lifecycle.name);
        // Having the environment made the graph too complex
        // Application Instances have environment prefix in their name
        // dataSnapShot.nodes.push({id: env.name, color: 'orange'});
        // dataSnapShot.links.push({source: lifecycleName, target: env.name});
        dataSnapShot.nodes.push({id: appInstance.name, color: 'green'});
        // Add lifecycle link
        dataSnapShot.links.push({source: lifecycleName, target: appInstance.name});
        // Add chart link
        const chartName = this.props.chartVersions[appInstance.chart_version_id].repo +
          ':' + this.props.chartVersions[appInstance.chart_version_id].version;
        if (chartName && chartName.length > 0) {
          dataSnapShot.links.push({source: appInstance.name, target: chartName});
        }
        this.setState({data: dataSnapShot});
      }
    });
  }

// Callback to handle click on the graph.
// @param {Object} event click dom event
  onClickGraph(event) {
    // window.alert('Clicked the graph background');
  }

  onClickNode(nodeId, node) {
    // window.alert('Clicked node ${nodeId} in position (${node.x}, ${node.y})');
  }

  onDoubleClickNode(nodeId, node) {
    // window.alert('Double clicked node ${nodeId} in position (${node.x}, ${node.y})');
  }

  onRightClickNode(event, nodeId, node) {
    // window.alert('Right clicked node ${nodeId} in position (${node.x}, ${node.y})');
  };

  onMouseOverNode(nodeId, node) {
    // console.info(`Mouse over node ${nodeId} in position (${node.x}, ${node.y})`);
  }

  onMouseOutNode(nodeId, node) {
    // window.alert(`Mouse out node ${nodeId} in position (${node.x}, ${node.y})`);
  }

  onClickLink(source, target) {
    // window.alert(`Clicked link between ${source} and ${target}`);
  }

  onRightClickLink(event, source, target) {
    // window.alert('Right clicked link between ${source} and ${target}');
  }

  onMouseOverLink(source, target) {
    // console.info(`Mouse over in link between ${source} and ${target}`);
  }

  onMouseOutLink(source, target) {
    // window.alert(`Mouse out link between ${source} and ${target}`);
  }

  onNodePositionChange(nodeId, x, y) {
    // console.log(`Node ${nodeId} moved to new position x= ${x} y= ${y}`);
  }

// Callback that's called whenever the graph is zoomed in/out
// @param {number} previousZoom the previous graph zoom
// @param {number} newZoom the new graph zoom
  onZoomChange(previousZoom, newZoom) {
    // window.alert(`Graph is now zoomed at ${newZoom} from ${previousZoom}`);
  }

  render() {
    return (
      <>
        <Graph
          id='graph-id' // id is mandatory, if no id is defined rd3g will throw an error
          data={this.state.data}
          config={this.state.config}
          onClickGraph={this.onClickGraph}
          onClickNode={this.onClickNode}
          onDoubleClickNode={this.onDoubleClickNode}
          onRightClickNode={this.onRightClickNode}
          onClickLink={this.onClickLink}
          onRightClickLink={this.onRightClickLink}
          onMouseOverNode={this.onMouseOverNode}
          onMouseOutNode={this.onMouseOutNode}
          onMouseOverLink={this.onMouseOverLink}
          onMouseOutLink={this.onMouseOutLink}
          onNodePositionChange={this.onNodePositionChange}
          onZoomChange={this.onZoomChange}/>
      </>
    );
  }

}

const mapStateToProps = state => {
  return {
    graph: state.graph,
  };
};

export default connect(mapStateToProps, {})(AppGraph);

