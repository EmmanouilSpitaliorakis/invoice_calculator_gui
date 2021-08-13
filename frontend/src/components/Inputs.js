import React from 'react'

function Inputs() {
    return (
        <form className="add-form">
            <div className="form-control">
                <label>Please enter the Applicant's ID</label>
                <input type="text" placeholder="Applicant ID"></input>
            </div>
            <div className="form-control">
                <label>Please enter the fees amount paid</label>
                <input type="text" placeholder="Fees Amount"></input>
            </div>
            <div className="form-control">
                <label>Please enter the tenancy length</label>
                <input type="text" placeholder="Tenancy in months"></input>
            </div>
            <div className="form-control form-control-radio">
                <label>Payment Type</label>
                <input type="radio" value="Full" name="payment"/>Full
                <input type="radio" value="Payment Plan" name="payment"/>Payment Plan
            </div>
            <div className="btn-row">
                <div className="btn-col-md-4 "><input type="submit" value="Calculate" className="btn"/></div>
                <div className="btn-col-md-4 "><input type="submit" value="Reset" className="btn"/></div>
                <div className="btn-col-md-4 "><input type="submit" value="Exit" className="btn"/></div>
            </div>
        </form>
    )
}


export default Inputs

