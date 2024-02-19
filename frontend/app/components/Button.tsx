'use client'

import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCalculator } from '@fortawesome/free-solid-svg-icons';
import { useNumber } from '../context/numbers';

const Button: React.FC = ({onClick}) => {
    const {number} = useNumber()

    return (
        <button
            className="flex flex-row items-center px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-300"
            type="button"
            onClick={onClick}
            disabled={number === ''}
        >
            <FontAwesomeIcon icon={faCalculator} className="mr-2" />
            <span>Index</span>
        </button>
    );
};

export default Button;
