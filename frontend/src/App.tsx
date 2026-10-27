import { useState } from 'react'
import './App.css'
import Navbar from './components/Navbar/Navbar'
import Proxy from './components/Proxy/Proxy'
import { Navbar_classes } from './components/Navbar/Nav_classes,'

function App() {
  let [mode,setmode]=useState<string>("proxy")
  return (
    <div>
      <Navbar setmode={setmode} classes={Navbar_classes} />
      {mode === "proxy" && <Proxy />}
      {/* {mode === "tester" && <Tester />} */}
    </div>
  );
}

export default App;

