import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTag } from '@fortawesome/free-solid-svg-icons';


interface ResultProps {
  number: number;
}

const Result: React.FC<ResultProps> = ({ number }) => {
  return (
    <div className="relative  flex items-center justify-center">
            <FontAwesomeIcon icon={faTag} size="2x" className="mr-2 absolute top-[10%] right-0 text-amber-600" />

      <div className="bg-white rounded-lg border border-gray-200 shadow-md p-6 m-4">
        <div className="flex flex-col items-center">
          <h3 className="text-lg leading-6 font-medium text-gray-900">
            Index
          </h3>
          <div className="mt-2 text-3xl font-semibold text-indigo-600">
            {number !== -1 ? number : 'Not Found'}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Result;
