import React from 'react';
import AddDevice from './components/AddDevice';
import GrafanaDashboard from './components/GrafanaDashboard';

function App() {
    return (
        <div>
            <h1>Network Monitoring</h1>
            <AddDevice />
            <GrafanaDashboard />
        </div>
    );
}

export default App;
