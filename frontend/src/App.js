import React from 'react';
import Header from './components/Header';
import Inputs from './components/Inputs';
import Button from "./components/Button";
import Output from "./components/Output";

function App() {

  const calculateOnClick = () => {
    console.log("Calculate")
  }

  const exitOnClick = () => {
    console.log("Exit")
  }

  return (
    <div id="app" className="App">
      <Header text="Housing Hand Invoice Calculator"/>
      <Inputs/>
      <Output/>
      <div className="btn-row">
        <div className="btn-col-md-4 "><Button text="Calculate" onClick = {calculateOnClick}/></div>
        <div className="btn-col-md-4 "><Button text="Exit" onClick={exitOnClick}/></div>
      </div>
    </div>
  );  
  }

export default App;
