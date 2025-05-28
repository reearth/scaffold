/* eslint-disable */
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Cursor: { input: any; output: any; }
  Time: { input: any; output: any; }
};

export type AddWorkspaceMemberInput = {
  role: Role;
  userId: Scalars['ID']['input'];
  workspaceId: Scalars['ID']['input'];
};

export type CreateProjectInput = {
  name: Scalars['String']['input'];
  workspaceID: Scalars['ID']['input'];
};

export type CreateTodoInput = {
  name: Scalars['String']['input'];
  projectId: Scalars['ID']['input'];
};

export type CreateWorkspaceInput = {
  name: Scalars['String']['input'];
};

export type Me = Node & {
  __typename?: 'Me';
  email: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  workspaces?: Maybe<WorkspaceConnection>;
};

export type Mutation = {
  __typename?: 'Mutation';
  addWorkspaceMember: WorkspaceMember;
  createProject: Project;
  createTodo: Todo;
  createWorkspace: Workspace;
  deleteProject: Scalars['ID']['output'];
  deleteTodo: Scalars['ID']['output'];
  deleteWorkspace: Scalars['ID']['output'];
  removeWorkspaceMember: Scalars['ID']['output'];
  updateProject: Project;
  updateTodo: Todo;
  updateWorkspace: Workspace;
  updateWorkspaceMemberRole: WorkspaceMember;
};


export type MutationAddWorkspaceMemberArgs = {
  input: AddWorkspaceMemberInput;
};


export type MutationCreateProjectArgs = {
  input: CreateProjectInput;
};


export type MutationCreateTodoArgs = {
  input: CreateTodoInput;
};


export type MutationCreateWorkspaceArgs = {
  input: CreateWorkspaceInput;
};


export type MutationDeleteProjectArgs = {
  projectId: Scalars['ID']['input'];
};


export type MutationDeleteTodoArgs = {
  todoId: Scalars['ID']['input'];
};


export type MutationDeleteWorkspaceArgs = {
  workspaceId: Scalars['ID']['input'];
};


export type MutationRemoveWorkspaceMemberArgs = {
  userId: Scalars['ID']['input'];
  workspaceId: Scalars['ID']['input'];
};


export type MutationUpdateProjectArgs = {
  input: UpdateProjectInput;
};


export type MutationUpdateTodoArgs = {
  input: UpdateTodoInput;
};


export type MutationUpdateWorkspaceArgs = {
  input: UpdateWorkspaceInput;
};


export type MutationUpdateWorkspaceMemberRoleArgs = {
  input: UpdatwWorkspaceMemberInput;
};

export type Node = {
  id: Scalars['ID']['output'];
};

export type PageInfo = {
  __typename?: 'PageInfo';
  endCursor?: Maybe<Scalars['Cursor']['output']>;
  hasNextPage: Scalars['Boolean']['output'];
  hasPreviousPage: Scalars['Boolean']['output'];
  startCursor?: Maybe<Scalars['Cursor']['output']>;
};

export type Project = Node & {
  __typename?: 'Project';
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  todos: TodoConnection;
  workspace?: Maybe<Workspace>;
  workspaceID: Scalars['ID']['output'];
};


export type ProjectTodosArgs = {
  filter: TodoFilter;
};

export type ProjectConnection = {
  __typename?: 'ProjectConnection';
  edges: Array<ProjectEdge>;
  pageInfo: PageInfo;
};

export type ProjectEdge = {
  __typename?: 'ProjectEdge';
  cursor: Scalars['Cursor']['output'];
  node: Project;
};

export type ProjectFilter = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  workspaceID?: InputMaybe<Scalars['ID']['input']>;
};

export type Query = {
  __typename?: 'Query';
  Node?: Maybe<Node>;
  Nodes?: Maybe<Array<Maybe<Node>>>;
  me?: Maybe<Me>;
  projects: ProjectConnection;
  todos: TodoConnection;
  user?: Maybe<User>;
  workspaces?: Maybe<Array<Workspace>>;
};


export type QueryNodeArgs = {
  id: Scalars['ID']['input'];
};


export type QueryNodesArgs = {
  ids: Array<Scalars['ID']['input']>;
};


export type QueryProjectsArgs = {
  filter: ProjectFilter;
};


export type QueryTodosArgs = {
  filter: TodoFilter;
};


export type QueryUserArgs = {
  email: Scalars['String']['input'];
};

export enum Role {
  Admin = 'ADMIN',
  Member = 'MEMBER',
  Owner = 'OWNER'
}

export type Todo = Node & {
  __typename?: 'Todo';
  createdAt: Scalars['Time']['output'];
  done: Scalars['Boolean']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  project?: Maybe<Project>;
  projectId: Scalars['ID']['output'];
  updatedAt: Scalars['Time']['output'];
};

export type TodoConnection = {
  __typename?: 'TodoConnection';
  edges?: Maybe<Array<TodoEdge>>;
  pageInfo: PageInfo;
};

export type TodoEdge = {
  __typename?: 'TodoEdge';
  cursor: Scalars['String']['output'];
  node: Todo;
};

export type TodoFilter = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  projectId?: InputMaybe<Scalars['ID']['input']>;
};

export type UpdateProjectInput = {
  id: Scalars['ID']['input'];
  name?: InputMaybe<Scalars['String']['input']>;
};

export type UpdateTodoInput = {
  done?: InputMaybe<Scalars['Boolean']['input']>;
  id: Scalars['ID']['input'];
  name?: InputMaybe<Scalars['String']['input']>;
};

export type UpdateWorkspaceInput = {
  id: Scalars['ID']['input'];
  name?: InputMaybe<Scalars['String']['input']>;
};

export type UpdatwWorkspaceMemberInput = {
  role: Role;
  userId: Scalars['ID']['input'];
  workspaceId: Scalars['ID']['input'];
};

export type User = Node & {
  __typename?: 'User';
  email: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
};

export type Workspace = Node & {
  __typename?: 'Workspace';
  id: Scalars['ID']['output'];
  members: Array<WorkspaceMember>;
  name: Scalars['String']['output'];
  projects: ProjectConnection;
};


export type WorkspaceProjectsArgs = {
  filter: ProjectFilter;
};

export type WorkspaceConnection = {
  __typename?: 'WorkspaceConnection';
  edges: Array<WorkspaceEdge>;
  pageInfo: PageInfo;
};

export type WorkspaceEdge = {
  __typename?: 'WorkspaceEdge';
  cursor: Scalars['Cursor']['output'];
  node: Workspace;
};

export type WorkspaceMember = {
  __typename?: 'WorkspaceMember';
  role: Role;
  user?: Maybe<User>;
  userId: Scalars['ID']['output'];
};
