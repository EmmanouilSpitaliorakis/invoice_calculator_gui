import React, { useState } from 'react';


function Inputs() {
    const [id, setId] = useState("")
    const [fees, setFees] = useState("")
    const [length, setLength] = useState("")
    const [type, setType] = useState("Full")

    // const onSubmit = (e) =>{

    //     e.preventDefault()

    //     if (!id){
    //         alert("The Applicants ID is empty")
    //     }
    // }

    return (
        <form className="add-form">
            <div className="form-control">
                <label>Please enter the Applicant's ID</label>
                <input type="text" placeholder="Applicant ID"
                value={id}
                onChange={(e)=> setId(e.target.value)}
                ></input>
            </div>
            <div className="form-control">
                <label>Please enter the fees amount paid</label>
                <input type="text" placeholder="Fees Amount"
                value={fees}
                onChange={(e)=> setFees(e.target.value)}></input>
            </div>
            <div className="form-control">
                <label>Please enter the tenancy length</label>
                <input type="text" placeholder="Tenancy in months"
                value={length}
                onChange={(e)=> setLength(e.target.value)}></input>
            </div>
            <div className="form-control form-control-radio">
                <label>Payment Type</label>
                <input type="radio" value={type} name="payment" checked = {true}
                onChange={(e)=> setType(e.target.value)}/>Full
                <input type="radio" value={type} name="payment"
                onChange={(e)=> setType(e.target.value)}/>Payment Plan
            </div>
        </form>
    )
}


export default Inputs

