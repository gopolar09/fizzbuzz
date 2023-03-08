import React, { useState } from 'react';
import './App.scss';
import * as api from "./services/api";

function App() {
  const [count, setCount] = useState(1);
  const [message, setMessage] = useState("");
  const { getFizzBuzz } = api;

  const handlePush = async() => {
    const resp = await getFizzBuzz(count+1);
    setMessage(resp.message);
    setCount(count + 1);
  }

  return (
    <div className="App">
      <header className="App-header">
        <div className="App-count">
          Your count<br/>
          {count}
        </div>
        <button
          className="App-button"
          onClick={handlePush}
        >
          Push me!
        </button>
        <div className="App-message">{message}</div>
      </header>
    </div>
  );
}

export default App;
