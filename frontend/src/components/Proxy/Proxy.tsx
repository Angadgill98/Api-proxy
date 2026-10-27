import React, { useState } from 'react'
import Sidebar from './Sidebar/Sidebar'
import Dashboard from './Dashboard/Dashboard'
import { Sidebar_classes } from './Sidebar/Sidebar_classes';
import { Dashboard_classes } from './Dashboard/Dashboard_classes';

const Proxy = () => {
    let [current_app, setapp] = useState<Application | null>(null);
  return (
    <div className='proxy_container'>
        
        <div className='proxy_background'>
            <Sidebar set_current_app={setapp} classes={Sidebar_classes} />
            <Dashboard app={current_app} classes={Dashboard_classes} />
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

