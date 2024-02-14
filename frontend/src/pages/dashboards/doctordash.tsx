// pages/dashboard/doctor.tsx

import React, { useState } from 'react';

const patients = [
  { id: 1, name: 'John Doe' },
  { id: 2, name: 'Jane Smith' },
  { id: 3, name: 'Alice Johnson' },
  // Add more patients as needed
];

const DoctorDashboardPage: React.FC = () => {
  const [selectedPatient, setSelectedPatient] = useState('');

  return (
    <div>
      <h1>Doctor Dashboard</h1>
      <div className="mt-4">
        <label htmlFor="patient" className="block mb-2">Select patient:</label>
        <select 
          id="patient" 
          className="w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:border-blue-500" 
          value={selectedPatient} 
          onChange={(e) => setSelectedPatient(e.target.value)} 
        >
          <option value="">Select patient</option>
          {patients.map(patient => (
            <option key={patient.id} value={patient.name}>{patient.name}</option>
          ))}
        </select>
      </div>
      <div className="mt-4">
        {/* Scribble pad component will go here */}
        {/* You can add your scribble pad component here */}
      </div>
    </div>
  );
};

export default DoctorDashboardPage;
