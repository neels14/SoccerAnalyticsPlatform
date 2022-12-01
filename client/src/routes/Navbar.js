import { useState, useEffect, useRef} from 'react'
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import "./Navbar.css"
import { ButtonBase } from '@mui/material';
import { Link, useMatch, useResolvedPath } from "react-router-dom";
import Divider from '@mui/material/Divider';
import Box from '@mui/material/Box';

import ClickAwayListener from '@mui/material/ClickAwayListener';
import Grow from '@mui/material/Grow';
import Paper from '@mui/material/Paper';
import Popper from '@mui/material/Popper';
import MenuItem from '@mui/material/MenuItem';
import MenuList from '@mui/material/MenuList';

function Navbar() {
  return (
      <AppBar position="static">
        <Toolbar className='navbar'>
          <ButtonBase component={Link} to={'/'}>
            <Typography variant="h4" component="div">
              <b>Soccer Analytics</b>
            </Typography>
          </ButtonBase>
          <Box sx={{display: 'flex'}}>
            <NavbarButton to='/worldcups'>World Cups</NavbarButton>
            <NavbarButton to='/countries'>Countries</NavbarButton>
            <NavbarButton to='/players'>Players</NavbarButton>
            <NavbarButton to='/matches'>Matches</NavbarButton>
            <Divider orientation="vertical" flexItem sx={{mx: 2}}/>
            <FeaturesMenu/>
          </Box>
        </Toolbar>
      </AppBar>
  );
}

function NavbarButton({to, children, ...props}) {
  const resolvedPath = useResolvedPath(to)
  const isActive = useMatch({ path: resolvedPath.pathname, end: true})
  return (
    <Button component={Link} to={to} color={"inherit"} variant={isActive ? "outlined" : "text"} {...props}>
      {children}
    </Button>
  )
}

function FeaturesMenu() {
  const [open, setOpen] = useState(false);
  const anchorRef = useRef(null);

  const handleToggle = () => {
    setOpen((prevOpen) => !prevOpen);
  };

  const handleClose = (event) => {
    if (anchorRef.current && anchorRef.current.contains(event.target)) {
      return;
    }

    setOpen(false);
  };

  function handleListKeyDown(event) {
    if (event.key === 'Tab') {
      event.preventDefault();
      setOpen(false);
    } else if (event.key === 'Escape') {
      setOpen(false);
    }
  }

  // return focus to the button when we transitioned from !open -> open
  const prevOpen = useRef(open);
  useEffect(() => {
    if (prevOpen.current === true && open === false) {
      anchorRef.current.focus();
    }

    prevOpen.current = open;
  }, [open]);

  const isActive = useMatch({ path: "/features/:feature", end: true })

  const featureMenuItems = []

  for (var i = 1; i <= 6; i++) {
    featureMenuItems.push(<FeatureMenuItem key={i} featureNum={i} onClick={handleClose}/>);
  }

  return (
      <div>
        <Button
          ref={anchorRef}
          id="composition-button"
          aria-controls={open ? 'composition-menu' : undefined}
          aria-expanded={open ? 'true' : undefined}
          aria-haspopup="true"
          onClick={handleToggle}
          color="inherit"
          variant={isActive ? "outlined" : "text"}
        >
          Features
        </Button>
        <Popper
          open={open}
          anchorEl={anchorRef.current}
          role={undefined}
          placement="bottom-start"
          transition
        >
          {({ TransitionProps, placement }) => (
            <Grow
              {...TransitionProps}
              style={{
                transformOrigin:
                  placement === 'bottom-start' ? 'left top' : 'left bottom',
              }}
            >
              <Paper>
                <ClickAwayListener onClickAway={handleClose}>
                  <MenuList
                    autoFocusItem={open}
                    id="composition-menu"
                    aria-labelledby="composition-button"
                    onKeyDown={handleListKeyDown}
                    variant="selectedMenu"
                  >
                    {featureMenuItems}
                  </MenuList>
                </ClickAwayListener>
              </Paper>
            </Grow>
          )}
        </Popper>
      </div>
  );
}

function FeatureMenuItem({featureNum, ...props}) {
  const isActive = useMatch({ path: `/features/${featureNum}`, end: true})
  return (
    <MenuItem component={Link} to={`/features/${featureNum}`} selected={isActive ? true : false} {...props}>Feature {featureNum}</MenuItem>
  )
}

export default Navbar;
