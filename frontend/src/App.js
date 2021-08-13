import React from 'react';
import Header from './components/Header';
import Inputs from './components/Inputs';

function App() {

  return (
    <div id="app" className="App">
      <Header text="Housing Hand Invoice Calculator"/>
      <Inputs/>
      {/* <Outputs/> */}
    </div>
  );  
  }

export default App;
