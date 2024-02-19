'use client'

import React, { createContext, useContext, useCallback, useState } from 'react';
import { NUMBERS } from '../api';

interface NumberContextType {
  number: string;
  setNumber: (number: string) => void;
  getIndex: (number: string) => Promise<any>; 
}

const NumberContext = createContext<NumberContextType>({
  number: '',
  setNumber: () => {},
  getIndex: async () => false, 
});

export const useNumber = () => useContext(NumberContext);

export const NumberProvider: React.FC = ({ children }) => {
  const [number, setNumber] = useState('');

  const getIndex = useCallback(async (number: string) => {
    try {
      const response = await fetch(`${NUMBERS}/${number}`);
      const data = await response.json();
      
      return data; 
    } catch (error) {
      console.error('ERROR', error);

      return false; 
    }
  }, []);

  return (
    <NumberContext.Provider value={{ number, setNumber, getIndex }}>
      {children}
    </NumberContext.Provider>
  );
};
