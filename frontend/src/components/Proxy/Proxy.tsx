import React from 'react'
import Sidebar from './Sidebar/Sidebar'
import Dashboard from './Dashboard/Dashboard'

const Proxy = () => {
  return (
    <div className='proxy_container'>
        
        <div className='proxy_background'>
            <Sidebar/>
            <Dashboard/>
        </div>


    </div>
  )
}

export default Proxy