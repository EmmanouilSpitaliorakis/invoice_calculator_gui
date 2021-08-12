import React, { useState } from 'react';
import Modal from 'react-modal';
import Header from './Header';



function HelloWorld() {
	const [showModal, setShowModal] = useState(false);
	const [result, setResult] = useState(null);

	const handleOpenModal = () => {
		setShowModal(true);

		window.backend.go_string().then((result) => setResult(<Header/>));
	};

	const handleCloseModal = () => {
		setShowModal(false);
	};

	return (
		<div className="App">
			<button className="btn" onClick={() => handleOpenModal()} type="button">
				Enter
      </button>
			<Modal
				appElement={document.getElementById("app")}
				isOpen={showModal}
				contentLabel="Minimal Modal Example"
			>
				{result}
				<button className="btn" onClick={() => handleCloseModal()}>Return</button>
			</Modal>
		</div>
	);
}

export default HelloWorld;
