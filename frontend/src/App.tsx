import React from 'react';
import logo from './logo.svg';
import Canvas from '../src/components/canvas/canvas'
import LiveGollection from 'livegollection-client';
import './App.css';

function App() {
  const liveGoll = new LiveGollection('ws://localhost:8080/livegollection')
  return (
    <div className="App">
      <Canvas/>
    </div>
  );
}

export default App;
