import React from 'react'

type NavbarProps = {
  setmode: React.Dispatch<React.SetStateAction<string>>;
};

const Navbar = ({setmode}:NavbarProps) => {

  return (
    <div className='navbar_container'>
        <div className='navbar_background'>
            <button onClick={() => setmode("proxy")}>Proxy</button>
            <button onClick={() => setmode("tester")}>Tester</button>
        </div>
    </div>
  )
}

export default Navbar