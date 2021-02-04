import React, {useEffect, useState} from "react";
import { Graph } from 'react-d3-graph';

const lookupLifeCycleName = (name) => {
  switch (name) {
    case 'gov-dev':
    case 'gov-stage':
    case 'gov-prod':
      return 'gov';
  }
  return name;
}

export default function AppGraph(props) {

   const mynodes = { nodes: [
      {id: 'dev', color: 'red', x: 100, y: 200, rank: 0},
      {id: 'integration', color: 'red', x: 250, y: 200, rank: 1},
      {id: 'stage', color: 'red', x: 400, y: 200, rank: 2},
      {id: 'prod', color: 'red', x: 550, y: 200, rank: 3},
      {id: 'gov', color: 'red', x: 700, y: 200, rank: 4},
    ]};

   const mylinks = { links: [
       {source: 'dev', target: 'integration'},
       {source: 'integration', target: 'stage'},
       {source: 'stage', target: 'prod'},
       {source: 'prod', target: 'gov'},
     ]};

   const [data, setData] = useState({nodes: mynodes.nodes, links: mylinks.links});

// graph payload (with minimalist structure)
//   const data = {
//     // directed: true,
//     nodes: [
//       {id: 'dev', color: 'red', cx: 350, cy: 750},
//       {id: 'integration', color: 'red', cx: 550, cy: 300},
//       {id: 'prod', color: 'red', cx: 300, cy: 300},
//
//       {id: 'env-1/sample'},
//       {id: 'env-2/sample'},
//       {id: 'int-1/sample'},
//       {id: 'int-2/sample'},
//       {id: 'prod-1/sample'},
//       {id: 'prod-2/sample'},
//       {id: 'chart.v1', color: 'gray'},
//       {id: 'chart.v2', color: 'gray'},
//       {id: 'chart.v3', color: 'gray'},
//       {id: 'chart.v4', color: 'gray'},
//     ],
//     links: [
//       {source: 'dev',
//         target: 'integration'
//       },
//       {source: 'integration', target: 'production'},
//       {source: 'env-1/sample', target: 'dev'},
//       {source: 'chart.v4', target: 'env-1/sample'},
//       {source: 'env-2/sample', target: 'dev'},
//       {source: 'chart.v3', target: 'env-2/sample'},
//       {source: 'chart.v3', target: 'int-1/sample'},
//       {source: 'int-1/sample', target: 'integration'},
//       {source: 'chart.v2', target: 'int-2/sample'},
//       {source: 'int-2/sample', target: 'integration'},
//       {source: 'chart.v2', target: 'prod-1/sample'},
//       {source: 'prod-1/sample', target: 'production'},
//       {source: 'chart.v1', target: 'prod-2/sample'},
//       {source: 'prod-2/sample', target: 'production'},
//
//     ]
//   };

// The graph configuration, you only need to pass down properties
// that you want to override, otherwise default ones will be used
  const myConfig = {
    // height: 400,
    // width: 800,
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

//   const myConfig = {
//     // 'automaticRearrangeAfterDropNode': false,
//     'collapsible': false,
//     'directed': true,
//     'focusAnimationDuration': 0.75,
//     'height': 700,
//     'highlightDegree': 2,
//     'highlightOpacity': 0.2,
//     'linkHighlightBehavior': true,
//     'maxZoom': 12,
//     'minZoom': 0.05,
//     'nodeHighlightBehavior': true,
//     'panAndZoom': false,
//     'staticGraph': false,
//     'staticGraphWithDragAndDrop': true,
//     'd3': {
//       'alphaTarget': 0.05,
//       'gravity': -100,
//       'linkLength': 150,
//       'linkStrength': 1
//     },
//     'node': {
//       'mouseCursor': 'grab',
//       'opacity': 1,
//       'renderLabel': true,
//       'size': 1500,
//       // 'viewGenerator': node => <QuestionNodeSvg question={node}/>
//     },
//     'link': {
//       'color': '#000',
//       'highlightColor': 'red',
//       'mouseCursor': 'pointer',
//       'opacity': 1,
//       'semanticStrokeWidth': false,
//       'strokeWidth': 2
//     }
//   };

  useEffect(() => {
    let dataSnapShot = data;
    // Add the charts to Graph
    let chartMap = {}
    props.applicationInstances.forEach((appInstance) => {
      chartMap[appInstance.chart_version_id] = appInstance.chart_version_id;
    });
    const chartIds = Object.values(chartMap)
    console.log(chartIds)
    chartIds.forEach(id => {
      const chartName = props.chartVersions[id].repo + ':' + props.chartVersions[id].version;
      console.log(chartName)
      if (chartName && chartName.length > 0) {
        dataSnapShot.nodes.push({id: chartName, color: 'gray'});
      };
    });

    props.applicationInstances.forEach((appInstance) => {
      const env = props.environments[appInstance.environment_id];
      console.log(env)
      const lifecycle = props.lifecycles[env.lifecycle_id];
      console.log(lifecycle)
      const lifecycleName = lookupLifeCycleName(lifecycle.name);
      console.log(lifecycleName)
      // Having the environment made the graph too complex
      // Application Instances have environment prefix in their name
      // dataSnapShot.nodes.push({id: env.name, color: 'orange'});
      // dataSnapShot.links.push({source: lifecycleName, target: env.name});
      dataSnapShot.nodes.push({id: appInstance.name, color: 'green'});
      // Add lifecycle link
      dataSnapShot.links.push({source: lifecycleName, target: appInstance.name});
      // Add chart link
      const chartName = props.chartVersions[appInstance.chart_version_id].repo +
        ':' + props.chartVersions[appInstance.chart_version_id].version;
      if (chartName && chartName.length > 0) {
        dataSnapShot.links.push({source: appInstance.name, target: chartName});
      };
      setData(dataSnapShot);
    });
  }, [props.applicationInstances, props.environments, props.lifecycles]);

  console.log("data....", data)
// Callback to handle click on the graph.
// @param {Object} event click dom event
  const onClickGraph = function(event) {
    // window.alert('Clicked the graph background');
  };

  const onClickNode = function(nodeId, node) {
    // window.alert('Clicked node ${nodeId} in position (${node.x}, ${node.y})');
  };

  const onDoubleClickNode = function(nodeId, node) {
    // window.alert('Double clicked node ${nodeId} in position (${node.x}, ${node.y})');
  };

  const onRightClickNode = function(event, nodeId, node) {
    // window.alert('Right clicked node ${nodeId} in position (${node.x}, ${node.y})');
  };

  const onMouseOverNode = function(nodeId, node) {
    console.info(`Mouse over node ${nodeId} in position (${node.x}, ${node.y})`);
  };

  const onMouseOutNode = function(nodeId, node) {
    // window.alert(`Mouse out node ${nodeId} in position (${node.x}, ${node.y})`);
  };

  const onClickLink = function(source, target) {
    // window.alert(`Clicked link between ${source} and ${target}`);
  };

  const onRightClickLink = function(event, source, target) {
    // window.alert('Right clicked link between ${source} and ${target}');
  };

  const onMouseOverLink = function(source, target) {
    // console.info(`Mouse over in link between ${source} and ${target}`);
  };

  const onMouseOutLink = function(source, target) {
    // window.alert(`Mouse out link between ${source} and ${target}`);
  };

  const onNodePositionChange = function(nodeId, x, y) {
    // console.log(`Node ${nodeId} moved to new position x= ${x} y= ${y}`);
  };

// Callback that's called whenever the graph is zoomed in/out
// @param {number} previousZoom the previous graph zoom
// @param {number} newZoom the new graph zoom
  const onZoomChange = function(previousZoom, newZoom) {
    // window.alert(`Graph is now zoomed at ${newZoom} from ${previousZoom}`);
  };

  return (
    <>
      <Graph
        id='graph-id' // id is mandatory, if no id is defined rd3g will throw an error
        data={data}
        config={myConfig}
        onClickGraph={onClickGraph}
        onClickNode={onClickNode}
        onDoubleClickNode={onDoubleClickNode}
        onRightClickNode={onRightClickNode}
        onClickLink={onClickLink}
        onRightClickLink={onRightClickLink}
        onMouseOverNode={onMouseOverNode}
        onMouseOutNode={onMouseOutNode}
        onMouseOverLink={onMouseOverLink}
        onMouseOutLink={onMouseOutLink}
        onNodePositionChange={onNodePositionChange}
        onZoomChange={onZoomChange}/>
    </>
  );
}
