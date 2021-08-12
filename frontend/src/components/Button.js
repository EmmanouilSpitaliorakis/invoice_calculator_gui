import React from 'react';
import PropTypes from 'prop-types';

const Button = ({ colour, text }) => {
    return (
        <button style={{ backgroundColor: colour }} className="btn" >{ text }</button>
    )
}


Button.defaultProps={
    colour: "steelblue"
}

Button.PropTypes={
    text: PropTypes.string,
    colour: PropTypes.string
}

export default Buttons
