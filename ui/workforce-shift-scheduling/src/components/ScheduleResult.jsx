import { useTranslation } from 'react-i18next';

export default function ScheduleResult({ result }) {
  const {t} = useTranslation();
  return (
    <div>
      <h2>{t('assigned_shifts')}</h2>
      {result.map((shift, idx) => (
        <div key={idx}>
          <strong>{shift.date} ({shift.slot})</strong>
          <ul>
            {shift.assigned.map((w) => (
              <li key={w.id}>{w.name}</li>
            ))}
          </ul>
        </div>
      ))}
    </div>
  );
}
