import React from 'react';
import type { NavbarClasses } from './Nav_classes,';

type NavbarProps = {
  setmode: React.Dispatch<React.SetStateAction<string>>;
  classes: NavbarClasses
};

const Navbar = ({ setmode, classes }: NavbarProps) => {
  return (
    <div className={classes.container}>
        <button className={classes.button} onClick={() => setmode("proxy")}>Proxy</button>
        <button className={classes.button} onClick={() => setmode("tester")}>Tester</button>
    </div>
  );
};

export default Navbar;