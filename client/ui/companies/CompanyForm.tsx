import { useState } from 'react';

import { Company } from '@/lib/types';
import Button from '@/ui/Button';

import { useFormStatus, useFormState } from 'react-dom';

import style from '@/styles/company/CompanyForm.module.css';

type message = null | string;
type state = {
  message: message;
};

export default function CompanyForm({
  action,
  onClose,
  company,
}: {
  // action is a function that must return a message
  action: (formData: FormData) => Promise<message>;
  onClose: () => void;
  company?: Company;
}) {
  const { pending } = useFormStatus();

  const [name, setName] = useState(company?.name);
  const [candidate_portal, setCandidate_portal] = useState(company?.candidate_portal);

  const [state, formAction] = useFormState<state, FormData>(
    async (state: state, formData: FormData): Promise<state> => {
      return {
        message: await action(formData),
      };
    },

    {
      message: null as message,
    },
  );

  return (
    <div>
      <form action={formAction}>
        <div className={style.form}>
          <label htmlFor="name">Name</label>
          <input
            required
            type="text"
            id="name"
            name="name"
            placeholder="Name"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          <label htmlFor="candidate_portal">Candidate Portal</label>
          <input
            required
            type="url"
            id="candidate_portal"
            name="candidate_portal"
            placeholder="Candidate portal"
            value={candidate_portal}
            onChange={(e) => setCandidate_portal(e.target.value)}
          />
        </div>
        <div className={style.btns}>
          <Button className={style.cancel} onClick={onClose}>
            Close
          </Button>
          <Button type="submit" disabled={pending}>
            Add
          </Button>
          {/* it should be the state.message here that would be returned from the server actions this is only for screen readers*/}
          <p aria-live="polite" className="sr-only" role="status">
            {state?.message}
          </p>
        </div>
      </form>
    </div>
  );
}
