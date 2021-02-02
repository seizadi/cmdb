import React from 'react';
import { connect } from 'react-redux';
import { makeStyles, withStyles } from '@material-ui/core/styles';
import InputLabel from '@material-ui/core/InputLabel';
import FormControl from '@material-ui/core/FormControl';
import NativeSelect from '@material-ui/core/NativeSelect';
import InputBase from '@material-ui/core/InputBase';

// from Redux
import { listEnvironments } from "../../actions";

const BootstrapInput = withStyles((theme) => ({
  root: {
    'label + &': {
      marginTop: theme.spacing(3),
    },
  },
  input: {
    borderRadius: 4,
    position: 'relative',
    backgroundColor: theme.palette.background.paper,
    border: '1px solid #ced4da',
    fontSize: 16,
    padding: '10px 26px 10px 12px',
    transition: theme.transitions.create(['border-color', 'box-shadow']),
    // Use the system font instead of the default Roboto font.
    fontFamily: [
      '-apple-system',
      'BlinkMacSystemFont',
      '"Segoe UI"',
      'Roboto',
      '"Helvetica Neue"',
      'Arial',
      'sans-serif',
      '"Apple Color Emoji"',
      '"Segoe UI Emoji"',
      '"Segoe UI Symbol"',
    ].join(','),
    '&:focus': {
      borderRadius: 4,
      borderColor: '#80bdff',
      boxShadow: '0 0 0 0.2rem rgba(0,123,255,.25)',
    },
  },
}))(InputBase);

const useStyles = makeStyles((theme) => ({
  margin: {
    margin: theme.spacing(1),
  },
}));

function EnvironmentSelects(props) {
  const classes = useStyles();
  const handleChange = (event) => {
    props.selectEnvironment(event.target.value);
  };

  React.useEffect(() => {
    // Specify how to clean up after this effect:
    props.listEnvironments();
  }, []);

  const getEnvOptions = () => {
    return (props.environments.map( env => {
        return <option key={env.id} value={env.id}>{env.name}</option>;
      }));
    };

  return (
    <div>
      <FormControl className={classes.margin}>
        <InputLabel htmlFor="environment-select-native">Environment</InputLabel>
        <NativeSelect
          id="environment-select-native"
          value={props.envId}
          onChange={handleChange}
          input={<BootstrapInput />}
        >
          {/*<option aria-label="None" value="" />*/}
          { getEnvOptions() }
        </NativeSelect>
      </FormControl>
    </div>
  );
}

function mapStateToProps(state) {
  return {
    environments: Object.values(state.environments),
  };
}

export default connect(mapStateToProps, { listEnvironments })(EnvironmentSelects);

