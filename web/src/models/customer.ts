/** The Customer metadata returned from the backend. */
export interface CustomerMeta {
  id: number;
  contact: {
    firstname: string;
    lastname: string;
    email: string;
    phone?: string;
  };
}
