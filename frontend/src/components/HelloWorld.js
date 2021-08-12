import React, { useState } from 'react';
import Modal from 'react-modal';

function HelloWorld() {
	const [showModal, setShowModal] = useState(false);
	const [result, setResult] = useState(null);

	const handleOpenModal = () => {
		setShowModal(true);

		window.backend.basic().then((result) => setResult(result));
	};

	const handleCloseModal = () => {
		setShowModal(false);
	};

	return (
		<div className="App">
			<button onClick={() => handleOpenModal()} type="button">
				Enter
      </button>
			<Modal
				appElement={document.getElementById("app")}
				isOpen={showModal}
				contentLabel="Minimal Modal Example"
			>
				<p>{result}</p>
				<button onClick={() => handleCloseModal()}>Return to logo</button>
			</Modal>
		</div>
	);
}

export default HelloWorld;
