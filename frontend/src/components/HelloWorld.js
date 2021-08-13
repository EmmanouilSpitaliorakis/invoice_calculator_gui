import React, { useState } from 'react';
import Modal from 'react-modal';
import Button from './Button';
import Header from './Header';



function HelloWorld() {
	const [showModal, setShowModal] = useState(false);
	const [result, setResult] = useState(null);

	const handleOpenModal = () => {
		setShowModal(true);

		window.backend.go_string().then((result) => setResult(<Header text="Invoice Calculator"/>));
	};

	const handleCloseModal = () => {
		setShowModal(false);
	};

	return (
		<div className="App">
			<Header text="Welcome to the Housing Hand Invoice Calculator"/>
			<Button text='Enter' onClick={handleOpenModal}/>
			<Modal
				appElement={document.getElementById("app")}
				isOpen={showModal}
				contentLabel="Minimal Modal Example"
			>
				{result}
				<Button text="Return" onClick={handleCloseModal}/>
			</Modal>
		</div>
	);
}

export default HelloWorld;
