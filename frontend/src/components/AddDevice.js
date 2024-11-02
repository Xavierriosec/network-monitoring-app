import React, { useState } from 'react';
import axios from 'axios';

function AddDevice() {
    const [device, setDevice] = useState({ name: '', ip: '' });

    const handleSubmit = async (e) => {
        e.preventDefault();
        await axios.post('/api/devices', device);
        alert("Device added!");
    };

    return (
        <form onSubmit={handleSubmit}>
            <input type="text" placeholder="Name" onChange={e => setDevice({...device, name: e.target.value})} />
            <input type="text" placeholder="IP" onChange={e => setDevice({...device, ip: e.target.value})} />
            <button type="submit">Add Device</button>
        </form>
    );
}

export default AddDevice;
