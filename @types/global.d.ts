declare function add(x: string, y: string): string;
declare function add(x: number, y: number): number;

declare type Person<T extends string | number> = {
	id: T;
	firstName: string;
	lastName?: string;
};
