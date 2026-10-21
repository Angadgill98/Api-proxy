import React, { useState } from 'react'
import Sidebar from './Sidebar/Sidebar'
import Dashboard from './Dashboard/Dashboard'

const Proxy = () => {
    let [current_app, setapp] = useState<Application | null>(null);
  return (
    <div className='proxy_container'>
        
        <div className='proxy_background'>
            <Sidebar set_current_app={setapp} />
            <Dashboard app={current_app}/>
        </div>


    </div>
  )
}

export default Proxy


export type Application = {
  id: number;
  name: string;
  url: string;
};

