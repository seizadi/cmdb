import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';

const useStyles = makeStyles((theme) => ({
  root: {
    '& > *': {
      margin: theme.spacing(1),
    },
  },
}));

export default function AppButton(props) {
  const classes = useStyles();

  const handleClick = () => {
    if (props.onClick) {
      return props.onClick(props.app);
    }
  };

  return (
    <span className={classes.root}>
      <Button variant="contained" color="primary" onClick={handleClick}>
        {props.app.name}
      </Button>
    </span>
  );
}
