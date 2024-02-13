// pages/dashboard/receptionist.tsx

import React, { useState } from 'react';

const doctors = [
  { id: 1, name: 'Dr. Smith' },
  { id: 2, name: 'Dr. Johnson' },
  { id: 3, name: 'Dr. Williams' },
  // Add more doctors as needed
];

const ReceptionistDashboardPage: React.FC = () => {
  const [patientId, setPatientId] = useState('');
  const [issue, setIssue] = useState('');
  const [selectedDoctor, setSelectedDoctor] = useState('');

  const handleCreateAppointment = () => {
    // Add logic here to create the appointment
    console.log('Creating appointment...');
    console.log('Patient ID:', patientId);
    console.log('Issue:', issue);
    console.log('Doctor assigned:', selectedDoctor);
    // Reset form fields
    setPatientId('');
    setIssue('');
    setSelectedDoctor('');
  };

  return (
    <div>
      <h1>Receptionist Dashboard</h1>
      <div className="mt-4">
        <label htmlFor="patientId" className="block mb-2">Patient ID:</label>
        <input 
          type="number" 
          id="patientId" 
          className="w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:border-blue-500" 
          placeholder="Enter patient ID" 
          value={patientId} 
          onChange={(e) => setPatientId(e.target.value)} 
        />
      </div>
      <div className="mt-4">
        <label htmlFor="issue" className="block mb-2">Issue:</label>
        <input 
          type="text" 
          id="issue" 
          className="w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:border-blue-500" 
          placeholder="Enter issue" 
          value={issue} 
          onChange={(e) => setIssue(e.target.value)} 
        />
      </div>
      <div className="mt-4">
        <label htmlFor="doctor" className="block mb-2">Doctor assigned:</label>
        <select 
          id="doctor" 
          className="w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:border-blue-500" 
          value={selectedDoctor} 
          onChange={(e) => setSelectedDoctor(e.target.value)} 
        >
          <option value="">Select doctor</option>
          {doctors.map(doctor => (
            <option key={doctor.id} value={doctor.name}>{doctor.name}</option>
          ))}
        </select>
      </div>
      <div className="mt-4">
        <button 
          className="bg-blue-500 text-black px-4 py-2 rounded-lg hover:bg-blue-600 focus:outline-none"
          onClick={handleCreateAppointment}
        >
          Create Appointment
        </button>
      </div>
    </div>
  );
};

export default ReceptionistDashboardPage;
