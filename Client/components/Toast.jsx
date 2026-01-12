'use client';

import { useEffect } from 'react';
import { X, CheckCircle, XCircle } from 'lucide-react';

// # 'Toast' Component #
const Toast = ({ message, type, onClose }) => {
    useEffect(() => {
        const timer = setTimeout(() => {
            onClose();
        }, 5000);

        return () => clearTimeout(timer);
    }, [onClose]);

    return (
        <>
            <div className={`fixed top-5 right-5 z-50 flex items-start space-x-3 p-4 rounded-lg shadow-lg max-w-md animate-in slide-in-from-right ${type === 'success' ? 'bg-green-100 border border-green-500' : 'bg-red-100 border border-red-500'
                }`}>
                {type === 'success' ? (
                    <CheckCircle className="w-5 h-5 text-green-600 flex-shrink-0 mt-0.5 font-medium" />
                ) : (
                    <XCircle className="w-5 h-5 text-red-600 flex-shrink-0 mt-0.5 font-medium" />
                )}
                <p className={`text-md flex-1 font-medium ${type === 'success' ? 'text-green-700' : 'text-red-700'}`}>
                    {message}
                </p>
                <button
                    onClick={onClose}
                    className={`flex-shrink-0 mt-0.5 font-bold text-gray-700 hover:text-gray-900`}
                >
                    <X className="w-5 h-5" />
                </button>
            </div>
        </>
    );
};

export default Toast;