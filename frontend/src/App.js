import React from 'react';
import Header from './components/Header';
import Inputs from './components/Inputs';
import Button from "./components/Button";
// import Output from "./components/Output";

function App() {

  return (
    <div id="app" className="App">
      <Header text="Housing Hand Invoice Calculator"/>
      <Inputs/>
      {/* <Output/> */}
      <div className="btn-row">
        <div className="btn-col-md-4 "><Button text="Calculate" /></div>
        <div className="btn-col-md-4 "><Button text="Exit" /></div>
      </div>
    </div>
  );  
  }

export default App;
