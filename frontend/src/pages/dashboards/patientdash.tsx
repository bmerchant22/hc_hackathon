// pages/dashboard/patient.tsx

import React from 'react';

const bookings = [
  { id: 1, imageUrl: '/booking1.jpg', date: '2024-02-13' },
  { id: 2, imageUrl: '/booking2.jpg', date: '2024-02-12' },
  { id: 3, imageUrl: '/booking3.jpg', date: '2024-02-11' },
  // Add more booking data as needed
];

const PatientDashboardPage: React.FC = () => {
  return (
    <div>
      <h1>Patient Dashboard</h1>
      <div className="grid grid-cols-3 gap-4">
        {bookings.map(booking => (
          <div key={booking.id} className="bg-gray-200 p-4 rounded-lg">
            <img src={booking.imageUrl} alt={`Booking ${booking.id}`} className="w-full h-auto mb-2" />
            <p>Date: {booking.date}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default PatientDashboardPage;
