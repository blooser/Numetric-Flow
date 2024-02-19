'use client'

import React from 'react';
import { useNumber } from '../context/numbers';

const NumberInput: React.FC = (onChange) => {
    const {number, setNumber} = useNumber()

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const value = e.target.value;
        const regex = /^[0-9]*$/;

        if (value === '' || regex.test(value)) {
            setNumber(value);
        }
    };

    return (
        <div className="flex items-center justify-center">
            <input
                type="text"
                value={number}
                onChange={handleInputChange}
                className="form-input px-4 py-2 border  rounded-md focus:outline-none focus:ring focus:border-blue-300 transition duration-150 ease-in-out"
                placeholder="Provide number"
            />
        </div>
    );
};

export default NumberInput;
