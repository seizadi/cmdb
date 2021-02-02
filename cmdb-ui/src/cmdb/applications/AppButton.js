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

  return (
    <span className={classes.root}>
      <Button variant="contained" color="primary">
        {props.name}
      </Button>
    </span>
  );
}
