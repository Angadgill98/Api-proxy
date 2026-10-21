import { useState } from 'react'
import './App.css'
import Navbar from './components/Navbar/Navbar'
import Proxy from './components/Proxy/Proxy'

function App() {
  let [mode,setmode]=useState<string>("proxy")
  return (
    <div>
      <Navbar setmode={setmode}/>
      {mode === "proxy" && <Proxy />}
      {/* {mode === "tester" && <Tester />} */}
    </div>
  )
}

export default App
